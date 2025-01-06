import os

class Config:
    SECRET_KEY = os.environ.get('SECRET_KEY') or 'you-will-never-guess'
    SQLALCHEMY_DATABASE_URI = r'sqlite:///C:\Users\Anubhav Ranjan\OneDrive\Desktop\NexTurn\NexTurn\AnubhavRanjan-NexTurn-Program\M3_Python_API_Assignments\cs2_BookBuddy_App\bookbuddy\db\bookbuddy.db'
    SQLALCHEMY_TRACK_MODIFICATIONS = False