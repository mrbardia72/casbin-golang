# simple_utils.py - A tiny utility library

def reverse_string(text):
    """
    Return a new string with the characters of the input string in reverse order.
    
    Parameters:
        text (str): The string to be reversed.
    
    Returns:
        str: The reversed string.
    """
    return text[::-1]

def count_words(sentence):
    """
    Return the number of words in a sentence by splitting on whitespace.
    
    Parameters:
        sentence (str): The input string to count words from.
    
    Returns:
        int: The number of words in the sentence.
    """
    return len(sentence.split())

def celsius_to_fahrenheit(celsius):
    """
    Convert a temperature from Celsius to Fahrenheit.
    
    Parameters:
        celsius (float): Temperature value in degrees Celsius.
    
    Returns:
        float: Equivalent temperature in degrees Fahrenheit.
    """
    return (celsius * 9/5) + 32
