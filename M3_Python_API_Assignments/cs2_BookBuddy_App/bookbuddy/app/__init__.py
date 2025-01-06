from flask import Flask
from flask_sqlalchemy import SQLAlchemy
from config import Config


db = SQLAlchemy()

def create_app():
    app = Flask(__name__)
    app.config.from_object(Config)
    
    db.init_app(app)

    from .routes.auth_routes import auth_bp as index_blueprint
    app.register_blueprint(index_blueprint)

    from .routes.protected_routes import protected_bp as protected_blueprint
    app.register_blueprint(protected_blueprint)

    from .routes.book_routes import book_bp as book_blueprint
    app.register_blueprint(book_blueprint)

    from .routes.genre_routes import genre_bp as genre_blueprint
    app.register_blueprint(genre_blueprint)
    
    with app.app_context():
        db.create_all()
        from app.models import init_genre
        init_genre()

        return app
