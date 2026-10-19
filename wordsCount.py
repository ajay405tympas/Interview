
from collections import Counter
import re

def top_k_words(text:str, k:int) -> list[tuple[str, int]]:
	words = re.findall(r"[a-z']+",text.lower())
	word_counts = Counter(words)
	return Counter(word_counts).most_common(k)


text = """
Python is great. Python is easy to learn.
Python is powerful. Go is fast. Go is simple.
"""

result = top_k_words(text, 3)
print(result)