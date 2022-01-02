load("time.star", "time")

load("render.star", "render")

#33 - 126

DISPLAY = "Hello World"

def main():
    main = []
    for i in range(33, 126):
        num = int((126-33) * random()) +33 
        main.append(render.Text(content = chr(num), font = "tom-thumb"))

    letters = [[] for i in range(len(DISPLAY) +1)]

    for i in range(0, len(DISPLAY)):
        print("i: " + str(i))
        letters[i] = main[:]

    for i in range(100):
        for j in range(1, len(DISPLAY)):
            if i % j == 0:
                letters[j] = letters[j][-1:] + letters[j][:-1]
    
    layout = []

    for i in range(len(DISPLAY)):
        layout.append(render.Animation(children = letters[i])) 
    
    
    return render.Root(
        child = render.Box(
            render.Column(
                main_align = "space_evenly",
                cross_align = "center",
                expanded = True,
                children = [
                    render.Row(
                        children= layout
                    ),
                    render.Text(
                        content="Do it.",
                        font = "tom-thumb"
                    ),
                    render.Row(
                        children= layout
                    )
                ]
            )
            
            
        )

    )

def random():
    """Return a pseudo-random number in [0, 1)"""
    print(((time.now().nanosecond / 1000) % 1000) * .001)
    return time.now().nanosecond / 1000 % 1000 * .001
