def generate_data():
    for _ in range(1_000_000):
        yield "hello world"

with open("example.txt", "w") as file:
    for chunk in generate_data():
        file.write(chunk)


    
