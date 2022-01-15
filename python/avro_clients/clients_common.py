

schema_str = """
    {
        "namespace": "confluent.io.examples.serialization.avro",
        "name": "User",
        "type": "record",
        "fields": [
            {"name": "name", "type": "string"},
            {"name": "favorite_number", "type": "int"},
            {"name": "favorite_color", "type": "string"}
        ]
    }
    """

class User(object):
    """
    User record
    Args:
        name (str): User's name
        favorite_number (int): User's favorite number
        favorite_color (str): User's favorite color
    """
    def __init__(self, name, favorite_number, favorite_color):
        self.name = name
        self.favorite_number = favorite_number
        self.favorite_color = favorite_color

    def user_to_dict(user, ctx):
        """
        Returns a dict representation of a User instance for serialization.
        Args:
            user (User): User instance.
            ctx (SerializationContext): Metadata pertaining to the serialization
                operation.
        Returns:
            dict: Dict populated with user attributes to be serialized.
        """
        return dict(name=user.name,
                    favorite_number=user.favorite_number,
                    favorite_color=user.favorite_color)

    def dict_to_user(obj, ctx):
        """
        Converts object literal(dict) to a User instance.
        Args:
            obj (dict): Object literal(dict)
            ctx (SerializationContext): Metadata pertaining to the serialization
                operation.
        """
        if obj is None:
            return None

        return User(name=obj['name'],
                    favorite_number=obj['favorite_number'],
                    favorite_color=obj['favorite_color'])
