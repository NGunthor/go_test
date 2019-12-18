##Observer

The Observer pattern refers to behavioral patterns at the object level.

The Observer pattern defines a one-to-many relationship between objects so that when the state of one object changes, all objects that depend on it are notified and updated automatically.

The main participants in the pattern are Publishers (Publisher) and Subscribers (Observer).

**There are two ways to receive notifications from the publisher:**

1. Pull method: After receiving a notification from the publisher, the subscriber must go to the publisher and collect (pull) the data on his own.

2. Push method: The publisher does not notify the subscriber about data updates, but independently delivers (pushes) the data to the subscriber.

**There are Push method**