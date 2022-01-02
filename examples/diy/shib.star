load("encoding/base64.star", "base64")
load("render.star", "render")


SHIB_ICON = base64.decode("""
iVBORw0KGgoAAAANSUhEUgAAABEAAAARCAYAAAA7bUf6AAAAmElEQVQ4jbWTMQ7AMAgDWRn7536F//UH7hBoUzBtllrykIAuBikiLpgCppAF0V6YArsMe/HYFNldb4V48djGa+HrnPp4GtNvM0DEv5pC1t/FeD1kwQUSi/OIEOFjzbV52TQFA7G7V0jZQ95PC8nNwP06qQGokAfIFZAHaALQ5TIxQEnxlqZLQlN8jbUMyKAMyx9xSd0v/k0nZYcRvOF5rv0AAAAASUVORK5CYII=
""")

def main():
    return render.Root(
        child = render.Box(
            render.Image(src = SHIB_ICON)
        )
    ) 