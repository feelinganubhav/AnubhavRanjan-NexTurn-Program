from .book import Genre, Book
from .user import User
from .. import db

def init_genre():
    # Add default genres if not already present
    if not Genre.query.first():
        default_genres = ['Fiction', 'Non-Fiction', 'Mystery', 'Sci-Fi', 'Fantasy']
        for genre_name in default_genres:
            genre = Genre(name=genre_name)
            db.session.add(genre)
        db.session.commit()

