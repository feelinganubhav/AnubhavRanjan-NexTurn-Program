from datetime import datetime
from .. import db

class Book(db.Model):
    __tablename__ = 'books'

    id = db.Column(db.Integer, primary_key=True)
    title = db.Column(db.String(120), nullable=False)
    author = db.Column(db.String(80), nullable=False)
    published_year = db.Column(db.Integer, nullable=False)
    genre = db.Column(db.String(50), db.ForeignKey('genres.name'), nullable=False)
    created_at = db.Column(db.DateTime, default=datetime.utcnow)

    def to_dict(self):
        """Convert the book object to a dictionary."""
        return {
            "id": self.id,
            "title": self.title,
            "author": self.author,
            "published_year": self.published_year,
            "genre": self.genre
        }

class Genre(db.Model):
    __tablename__ = 'genres'

    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(50), unique=True, nullable=False)

    def to_dict(self):
        """Convert Genre object to a dictionary."""
        return {
            "id": self.id,
            "name": self.name
        }
