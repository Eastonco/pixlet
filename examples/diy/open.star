load("render.star", "render")


def main():
    return render.Root(
        child = render.Box(
            width = 64,
            height = 32,
            child = render.Text(
                content = "OPEN",
                color = "#ff0000",
                font = "CG-pixel-4x5-mono"
            )
        )
    )