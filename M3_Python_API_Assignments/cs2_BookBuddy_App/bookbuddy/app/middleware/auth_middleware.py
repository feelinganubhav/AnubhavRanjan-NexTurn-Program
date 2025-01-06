from flask import request, jsonify
import jwt
from functools import wraps
from config import Config
from ..models import User

SECRET_KEY = Config.SECRET_KEY

def token_required(f):
    @wraps(f)
    def decorated(*args, **kwargs):
        token = request.headers.get('Authorization')

        if not token:
            return jsonify({"error": "Token is missing"}), 401

        try:
            token = token.split(" ")[1]  # Bearer <token>
            decoded = jwt.decode(token, SECRET_KEY, algorithms=['HS256'])
            current_user = User.query.get(decoded['user_id'])
            if not current_user:
                raise Exception("User not found")
        except Exception as e:
            return jsonify({"error": "Invalid or expired token", "message": str(e)}), 401

        return f(current_user, *args, **kwargs)

    return decorated
