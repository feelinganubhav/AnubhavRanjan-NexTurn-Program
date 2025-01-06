from flask import Blueprint, request, jsonify
from sqlalchemy.exc import SQLAlchemyError
from ..models import Book, db
from ..middleware.auth_middleware import token_required

book_bp = Blueprint('books', __name__)

@book_bp.route('/books', methods=['GET'])
@token_required
def get_books(current_user):
    """Retrieve all books with pagination, sorting, and filtering by genre."""
    try:
        # Pagination
        page = int(request.args.get('page', 1))
        per_page = int(request.args.get('per_page', 10))
        
        # Sorting
        sort_by = request.args.get('sort_by', 'title')
        order = request.args.get('order', 'asc')
        order_by = getattr(Book, sort_by, Book.title)
        if order == 'desc':
            order_by = order_by.desc()
        
        # Filtering
        genre = request.args.get('genre')
        query = Book.query
        if genre:
            query = query.filter_by(genre=genre)
        
        # Pagination and Sorting
        books = query.order_by(order_by).paginate(page=page, per_page=per_page, error_out=False)
        data = [book.to_dict() for book in books.items]

        return jsonify({
            "books": data,
            "total": books.total,
            "pages": books.pages,
            "current_page": books.page
        }), 200
    except SQLAlchemyError as e:
        return jsonify({"error": "Database error", "message": str(e)}), 500


@book_bp.route('/books', methods=['POST'])
@token_required
def add_book(current_user):
    """Add a new book to the collection."""
    data = request.get_json()
    try:
        title = data['title']
        author = data['author']
        published_year = data['published_year']
        genre = data['genre']

        # Input validation
        if not title or not author or not published_year or not genre:
            return jsonify({"error": "Missing required fields"}), 400

        new_book = Book(title=title, author=author, published_year=published_year, genre=genre)
        db.session.add(new_book)
        db.session.commit()

        return jsonify({"message": "Book added successfully", "book": new_book.to_dict()}), 201
    except KeyError as e:
        return jsonify({"error": f"Missing field: {e.args[0]}"}), 400
    except SQLAlchemyError as e:
        return jsonify({"error": "Database error", "message": str(e)}), 500


@book_bp.route('/books/<int:book_id>', methods=['GET'])
@token_required
def get_book(current_user, book_id):
    """Retrieve a book by its ID."""
    book = Book.query.get(book_id)
    if not book:
        return jsonify({"error": "Book not found"}), 404
    return jsonify(book.to_dict()), 200


@book_bp.route('/books/<int:book_id>', methods=['PUT'])
@token_required
def update_book(current_user, book_id):
    """Update an existing book."""
    data = request.get_json()
    book = Book.query.get(book_id)
    if not book:
        return jsonify({"error": "Book not found"}), 404
    try:
        book.title = data.get('title', book.title)
        book.author = data.get('author', book.author)
        book.published_year = data.get('published_year', book.published_year)
        book.genre = data.get('genre', book.genre)
        
        db.session.commit()
        return jsonify({"message": "Book updated successfully", "book": book.to_dict()}), 200
    except SQLAlchemyError as e:
        return jsonify({"error": "Database error", "message": str(e)}), 500


@book_bp.route('/books/<int:book_id>', methods=['DELETE'])
@token_required
def delete_book(current_user, book_id):
    """Delete a book by its ID."""
    book = Book.query.get(book_id)
    if not book:
        return jsonify({"error": "Book not found"}), 404
    try:
        db.session.delete(book)
        db.session.commit()
        return jsonify({"message": f"Book with ID {book_id} deleted successfully"}), 200
    except SQLAlchemyError as e:
        return jsonify({"error": "Database error", "message": str(e)}), 500
