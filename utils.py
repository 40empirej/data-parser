import logging
import re
import json

# Set up logging
logger = logging.getLogger(__name__)

def is_valid_email(email: str) -> bool:
    """
    Validate an email address using a regular expression.
    
    :param email: The email address to validate
    :return: True if the email is valid, False otherwise
    """
    pattern = r"^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$"
    return bool(re.match(pattern, email))

def load_json(file_path: str) -> dict:
    """
    Load a JSON file and return its contents as a Python dictionary.
    
    :param file_path: The path to the JSON file
    :return: The contents of the JSON file as a dictionary
    """
    try:
        with open(file_path, 'r') as file:
            return json.load(file)
    except json.JSONDecodeError as e:
        logger.error(f"Error parsing JSON: {e}")
        return {}
except FileNotFoundError:
    logger.error(f"File not found: {file_path}")
    return {}