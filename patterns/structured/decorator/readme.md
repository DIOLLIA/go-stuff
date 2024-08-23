## Decorator (Wrapper)

Attach new behaviors to objects by placing these objects inside special wrapper objects that contain the behaviors.

“Wrapper” is the alternative nickname for the Decorator pattern that clearly expresses the main idea of the pattern. A wrapper is an object that can be linked with some target object. The wrapper contains the same set of methods as the target and delegates to it all requests it receives. However, the wrapper may alter the result by doing something either before or after it passes the request to the target.

Unlike the Proxy it can extend behaviour of the object.
Proxy grant the access to the original object, wrapper adds new behaviour.
Proxy usually used for access control, logging, caching. Decorator dynamically can extend the object