## Observer

is a behavioral design pattern that lets you define a subscription mechanism to notify multiple objects about any events that happen to the object they’re observing.

Conceptual Example:
In the website, items go out of stock from time to time. There can be customers who are interested in a particular item that went out of stock.

Best here would be that the customer subscribes only to the particular item he is interested in and gets notified if the item is available. Also, multiple customers can subscribe to the same product.
Option 3 is most viable, and this is what the Observer pattern is all about. The major components of the observer pattern are:

Subject, the instance which publishes an event when anything happens.
Observer, which subscribes to the subject events and gets notified when they happen.