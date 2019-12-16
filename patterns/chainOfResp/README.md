## Chain Of Responsibility

Chain Of Responsibility refers to behavioral patterns at the object level.

The Chain Of Responsibility pattern avoids the binding of the requesting object to the recipient of the request, while giving a chance to process this request to several objects. Recipients are linked in a chain, and the request is transmitted in a chain until it is processed by some object.

In essence, this is a chain of handlers that take turns receiving a request and then decide whether to process it or not. If the request is not processed, then it is passed down the chain. If it is processed, then the pattern itself decides to pass it on or not. If the request is not processed by any handler, then it is simply lost.

**Required to implement:**

1. The base abstract class `Handler`, describing the interface of handlers in chains;

2. The `concreteHandlerA` class that implements the concrete handler A;

3. The `concreteHandlerB` class that implements the concrete B handler;

4. The `concreteHandlerC` class that implements the concrete C handler;