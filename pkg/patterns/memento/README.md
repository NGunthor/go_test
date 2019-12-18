## **Memento**

The Memento pattern refers to behavioral patterns of the object level.

The Memento pattern receives and stores its internal state outside the object so that it can later be restored to the same state. If the client subsequently needs to "roll back" the state of the original object, it passes Memento back to the original object to restore it.

The pattern operates on three objects:

<ol>
<li>The owner of the state (Originator);</li>
<li>Keeper (Memento) - Stores the state of the object-owner of the Originator class;</li>
<li>Caretaker - Responsible for the safety of the custodian of the Memento class.</li>
</ol>

