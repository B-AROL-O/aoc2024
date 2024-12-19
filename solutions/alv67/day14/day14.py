# --- Part Two ---

# from icecream import ic
import re
import time
import random
from IPython.display import clear_output
import tkinter as tk

puzzle = {"filename": "input.txt", "width": 101, "height": 103}

# decomment next line for test
# puzzle = {"filename": "test.txt", "width": 11, "height": 7} #


class Robot:
    def __init__(self, x, y, vx, vy, room_instance):
        self.x = x
        self.y = y
        self.vx = vx
        self.vy = vy
        self.room = room_instance

    def position(self):
        return (self.x, self.y)

    def move(self, steps = 1):
        self.x += self.vx * steps  
        self.x = self.x % self.room.width
        self.y += self.vy * steps
        self.y = self.y % self.room.height

class Room:

    def __init__(self, width, height):    
        self.width = width
        self.height = height
        self.robots = []
        self.abyss = 0
        self.sand = 0
        self.time = 0

    def add_robot(self, x, y, vx, vy):
        robot_instance = Robot(x, y, vx, vy, self)
        self.robots.append(robot_instance)

    def time_step(self, seconds = 1):
        for robot in self.robots:
            robot.move(seconds)
        self.time += seconds

    def safety_factor(self):
        """ 
        count robots in each quadrant (robot exactly in the middle do not count)
        The safety factor is the multiplicationof those 4 counts
        """
        quad_w = self.width//2
        quad_h = self.height//2
        positions = [robot.position() for robot in self.robots]

        q1 = len([r for r in positions if r[0]<quad_w and r[1]<quad_h])
        q2 = len([r for r in positions if r[0]>quad_w and r[1]<quad_h])
        q3 = len([r for r in positions if r[0]>quad_w and r[1]>quad_h])
        q4 = len([r for r in positions if r[0]<quad_w and r[1]>quad_h])

        return q1 * q2 * q3 * q4

    def get_map(self):
        retmap = list()
        side2side = list()
        robot_positions = [robot.position() for robot in self.robots]
        for y in range(self.height):
            row = []
            s2s = 0
            s = 0
            for x in range(self.width):
                n = robot_positions.count((x, y))
                if n == 0:
                    if s > s2s:
                        s2s = s
                        s=0
                    row.append(".")
                else:
                    s+=1 # number of robots side to side
                    row.append(str(n))
            retmap.append(row)
            side2side.append(s2s)
        return retmap, max(side2side)


    
    def __str__(self):
        # extract robot positions
        retstr = ""
        robot_positions = [robot.position() for robot in self.robots]
        for y in range(self.height):
            for x in range(self.width):
                n = robot_positions.count((x, y))
                if n == 0:
                    retstr += "."
                else:
                    retstr += str(n)
            retstr += '\n'
        return retstr

# Dimensioni della mappa
MAP_WIDTH = 20
MAP_HEIGHT = 10

# Caratteri per la mappa
CHARS = ['#', '.', '*', ' ']
def generate_map():
    """Genera una nuova mappa casuale."""
    return [[random.choice(CHARS) for _ in range(MAP_WIDTH)] for _ in range(MAP_HEIGHT)]

def draw_map(canvas, char_map):
    """Disegna la mappa sul canvas."""
    canvas.delete("all")  # Pulisci il canvas
    for y, row in enumerate(char_map):
        for x, char in enumerate(row):
            canvas.create_text(
                x * 9,  # Posizione x
                y * 9,  # Posizione y
                text=char,
                font=("Courier", 8),
                anchor="nw"  # Inizia il testo dall'angolo in alto a sinistra
            )

# def update_map(canvas, root, get_map):
#     """Aggiorna la mappa e ridisegna."""
#     new_map = get_map()  # Genera una nuova mappa
#     draw_map(canvas, new_map)  # Disegna la mappa sul canvas
#     #root.after(500, update_map, canvas, root)  # Aggiorna dopo 500ms


def update_map(canvas, root, restroom):    
    side2side = 0
    new_map = []
    while side2side < 10:
        restroom.time_step(1)
        new_map, side2side = restroom.get_map()
        print (f"\r {restroom.time} - {side2side}       ", end="")
    print (f"\rtime: {restroom.time} side-2-side: {side2side}")
    draw_map(canvas, new_map)
    root.after(1000, update_map, canvas, root, restroom)  # Aggiorna dopo 500ms
# ---- main -----

pattern = r"p=(-?\d+),(-?\d+)\s+v=(-?\d+),(-?\d+)"

robot_list = []

with open(puzzle["filename"], "r") as file:

    for line in file:
        match = re.match(pattern, line.strip())
        if match:
            x, y, vx, vy = map(int, match.groups())
            robot_list.append((x, y, vx, vy))

restroom = Room(puzzle["width"], puzzle["height"])

for robot in robot_list:
    x,y,vx,vy = robot
    restroom.add_robot(x,y,vx,vy)



root = tk.Tk()
root.title("Restroom")

new_map = restroom.get_map()
canvas = tk.Canvas(root,width = puzzle["width"]*9, height = puzzle["height"]*9, bg="black" )
canvas.pack()


# start form...
restroom.time_step(7000)

update_map(canvas, root, restroom)

root.mainloop()

ret1 = restroom.safety_factor()

print (f"The solution of for part 1 is: {ret1}")
