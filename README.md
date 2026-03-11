# PROJECT SUMMARY

In this project, I built a command-line text processing tool using the Go programming language that reads text from an input file, applies several transformations, and prints the corrected version of the text. The program detects special editing commands such as (up), (low), (cap), and numbered variations like (up,2) to modify the previous word or group of words accordingly.

The application processes the text step by step by splitting it into words, identifying commands, and applying the required transformations while removing the command tokens from the final output. It also improves text quality by correcting punctuation spacing, adjusting grammar cases such as converting “a” to “an” before vowels, and handling edge cases that may appear in real text.

Through this project, I practiced file handling, loops, conditional logic, and string manipulation, while learning how to design modular functions that keep the code simple, readable, and reliable when processing different text scenarios.