## Adapter (Wrapper)

Allows objects with incompatible interfaces to collaborate.

An adapter wraps one of the objects to hide the complexity of conversion happening under the hood. The wrapped object isn’t aware of the adapter. 

The adapter gets an interface, compatible with one of the existing objects.
Using this interface, the existing object can safely call the adapter’s methods.
Upon receiving a call, the adapter passes the request to the second object, but in a format and order that the second object expects.