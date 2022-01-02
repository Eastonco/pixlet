load("encoding/base64.star", "base64")
load("render.star", "render")

# Turbo modifier, changes speed of car. Slowest is 1
# increase the range() on line 106 if encountering car overlap
TURBO = 2

# Background
TRACK = base64.decode("""iVBORw0KGgoAAAANSUhEUgAAAEAAAAAgCAYAAACinX6EAAAA3ElEQVRoge2YwQ2DMAxF3SG6E+LOMMzQsToDs3D8PSFVqHHcKPlxwJa+Ig6Wv+1HkJDXe8adJb0N9FYM4PthEwEA2nkEu25yAACwLhNNIkKtty4TNhGdAKYZ9sCPmkFAEBAE/EfAqF+BagSck0aQtswiAnIF2WHxQyVAi/35KNavfIsfNwS0aH4YAlLmraHluyfAss2SzQ9BQOvmXROQaygX1tfGJQGs5i1+ut8BrcMdAd5EJ8CbqhPA/m1V46xGwJVURMCVFAQEAUGAToCH27v1mRzAHRUD6G2gtz7RxEkVKU4XDAAAAABJRU5ErkJggg==""")

# Car Positions
R = base64.decode("iVBORw0KGgoAAAANSUhEUgAAAAQAAAADCAYAAAC09K7GAAAAHElEQVQImWNgYGD4r7nt2384rbnt239kzICuAgCHMRqdiFedjQAAAABJRU5ErkJggg==")
L = base64.decode("iVBORw0KGgoAAAANSUhEUgAAAAQAAAADCAYAAAC09K7GAAAAHUlEQVQImWPQ3PbtPwMDw38YzaC57dt/ZMyArgIApIEanVkREDsAAAAASUVORK5CYII=")
U = base64.decode("iVBORw0KGgoAAAANSUhEUgAAAAMAAAAECAYAAABLLYUHAAAAGUlEQVQImWNgXvnlPwwzMDAwwBj/GXDKAACgZBj1VrZDbQAAAABJRU5ErkJggg==")
D = base64.decode("iVBORw0KGgoAAAANSUhEUgAAAAMAAAAECAYAAABLLYUHAAAAH0lEQVQImWNgYGD4z7zyy38GBob/DMwrv/yHYQacMgBL5Bj1YlBw0gAAAABJRU5ErkJggg==")
RU = base64.decode("iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAJklEQVQImWNgQID/zCu//Gde+eU/XAQmABP8jybwH0UVAxYAFwQA+poeM27aO3MAAAAASUVORK5CYII=")
RD = base64.decode("iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAI0lEQVQImWNgQID/DMiAeeWX/8wrv6AI/ocJQiX+Y6jGqQMA1A4eMzkyzR4AAAAASUVORK5CYII=")
LU = base64.decode("iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAJ0lEQVQImWNgXvnlP/PKL/8ZGBj+M8AATBAqAQf/0SQwdTCgAbgAAGOBHjMBfB5jAAAAAElFTkSuQmCC")
LD = base64.decode("iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAJUlEQVQImWNgQID/DOiAeeWX/8wrv8Al/sMEoIL/GdAEMLTBBQE89R4zZHJwOAAAAABJRU5ErkJggg==")

def drawRace():
    path = []
    for i in range(0, 52, TURBO): 
        path.append(
            render.Box(
                child = render.Image(R),
                width = 58 + i,
                height = 12
            )
        )
    for i in range (0, 5, TURBO):
        path.append(
            render.Box(
                child = render.Image(RD),
                width = 110 + i,
                height = 12 + i
            )
        )
    for i in range(0, 28,TURBO): 
        path.append(
            render.Box(
                child = render.Image(D),
                width = 114,
                height = 20 + i
            )
        )
    for i in range (0, 5, TURBO):
        path.append(
            render.Box(
                child = render.Image(LD),
                width = 114 - i,
                height = 47 + i
            )
        )
    for i in range(0, 88,TURBO): 
        path.append(
            render.Box(
                child = render.Image(L),
                width = 108 - i,
                height = 50
            )
        )
    for i in range (0, 4,TURBO):
        path.append(
            render.Box(
                child = render.Image(LU),
                width = 16 - i,
                height = 50 - i
            )
        )
    for i in range(0, 26,TURBO): 
        path.append(
            render.Box(
                child = render.Image(U),
                width = 12,
                height = 45 - i
            )
        )
    for i in range (0, 6, TURBO):
        path.append(
            render.Box(
                child = render.Image(RU),
                width = 12 + i,
                height = 17 - i
            )
        )
    for i in range(0, 38,TURBO):
        path.append(
            render.Box(
                child = render.Image(R),
                width = 20 + i,
                height = 12
            )
        )
    return path

def main():
    racer1 = drawRace()
    racer2 = racer1[:]
    racer3 = racer1[:]
    racer4 = racer1[:]
    racer5 = racer1[:]
    racer6 = racer1[:]
    racer7 = racer1[:]
    racer8 = racer1[:]
    for i in range(800):
        if i % 2 == 0:
            racer2 = racer2[-1:] + racer2[:-1]
        if i % 3 == 0:
            racer3 = racer3[-1:] + racer3[:-1]
        if i % 4 == 0:
            racer4 = racer4[-1:] + racer4[:-1]
        if i % 5 == 0:
            racer5 = racer5[-1:] + racer5[:-1]
        if i % 6 == 0:
            racer6 = racer6[-1:] + racer6[:-1]
        if i % 7 == 0:
            racer7 = racer7[-1:] + racer7[:-1]
        if i % 8 == 0:
            racer8 = racer8[-1:] + racer8[:-1]


    return render.Root(
        child = render.Stack(
            children = [
            render.Image(src=TRACK),
            render.Animation(
                children = racer1
            ),
            render.Animation(
                children = racer2
            ),
            render.Animation(
                children = racer3
            ),
            render.Animation(
                children = racer4
            ),
            render.Animation(
                children = racer5
            ),
            render.Animation(
                children = racer6
            ),
            render.Animation(
                children = racer7
            ),
            render.Animation(
                children = racer8
            ),
                
            ]
        )
    )
