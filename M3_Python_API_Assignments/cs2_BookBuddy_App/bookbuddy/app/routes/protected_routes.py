from flask import Blueprint, jsonify
from ..middleware.auth_middleware import token_required

protected_bp = Blueprint('protected', __name__)

@protected_bp.route('/protected', methods=['GET'])
@token_required
def protected_route(current_user):
    return jsonify({
        "message": f"Welcome {current_user.username}! This is a protected route."
    }), 200
