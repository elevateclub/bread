# Story Event Presence System
measure observations and pass it through an algorithm that scores something
let's try to figure out the scoring system for our own lives remembering that numbers never lie but don't always tell the whole story.
so let's try to tell the whole story with seps: the Story Event Presence System.

Security is the appropriate monitoring of events that occur in time to trigger alerts when some inevitably tries to snoop in on your operation.
Trust is the agreement of terms between two parties.
A plan is established and executed on to deliver value.
Live our lives and tell the story.
Be present.
Let's take back ownership of our data.

## Learn to storyctl

Have Golang setup and install the tool 
```
go install https://github.com/breadtech/seps/cmd/storyctl
```

### Define storage, sensors, and securities
```
# Store: fs, obj, sqlite, mongo, postgres
# Sensors: collects events. Life comes with the basic, modify and define your own sensors.
# Securities: trigger alerts based on threshold conditions.
storyctl config --store fs --sensors life --securities air,water,food,sleep,activity,shelter
```

### Add events from life
```
# after you wake up
storyctl post event --type wakeup

# track consumption as an event
storyctl post consume water --amt 500mL

# workouts are many events
storyctl post workout --type opm

# before you sleep
storyctl post event --type sleep
``
