from flask import Blueprint, request, jsonify
from sqlalchemy.exc import SQLAlchemyError
from ..models import Genre, db
from ..middleware.auth_middleware import token_required

genre_bp = Blueprint('genres', __name__)

@genre_bp.route('/genres', methods=['GET'])
@token_required
def get_genres(current_user):
    """Retrieve all genres."""
    try:
        genres = Genre.query.all()
        return jsonify([genre.to_dict() for genre in genres]), 200
    except SQLAlchemyError as e:
        return jsonify({"error": "Database error", "message": str(e)}), 500


@genre_bp.route('/genres', methods=['POST'])
@token_required
def add_genre(current_user):
    """Add a new genre."""
    data = request.get_json()
    try:
        name = data['name']
        if not name:
            return jsonify({"error": "Genre name is required"}), 400

        new_genre = Genre(name=name)
        db.session.add(new_genre)
        db.session.commit()

        return jsonify({"message": "Genre added successfully", "genre": new_genre.to_dict()}), 201
    except KeyError as e:
        return jsonify({"error": f"Missing field: {e.args[0]}"}), 400
    except SQLAlchemyError as e:
        return jsonify({"error": "Database error", "message": str(e)}), 500


@genre_bp.route('/genres/<int:genre_id>', methods=['PUT'])
@token_required
def update_genre(current_user, genre_id):
    """Update an existing genre."""
    data = request.get_json()
    genre = Genre.query.get(genre_id)
    if not genre:
        return jsonify({"error": "Genre not found"}), 404
    try:
        genre.name = data.get('name', genre.name)
        db.session.commit()

        return jsonify({"message": "Genre updated successfully", "genre": genre.to_dict()}), 200
    except SQLAlchemyError as e:
        return jsonify({"error": "Database error", "message": str(e)}), 500


@genre_bp.route('/genres/<int:genre_id>', methods=['DELETE'])
@token_required
def delete_genre(current_user, genre_id):
    """Delete a genre by its ID."""
    genre = Genre.query.get(genre_id)
    if not genre:
        return jsonify({"error": "Genre not found"}), 404
    try:
        db.session.delete(genre)
        db.session.commit()

        return jsonify({"message": f"Genre with ID {genre_id} deleted successfully"}), 200
    except SQLAlchemyError as e:
        return jsonify({"error": "Database error", "message": str(e)}), 500
