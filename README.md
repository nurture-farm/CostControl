This library helps with monitoring and putting thresholds on any intended metrics. Primarily developed for monitoring monetory costs.

Pre-requisites:
Grafana setup: Used for setting up alerts
Prometheus setup: Used for storing the metrics

This library currently exposes 2 methods:
1) Configure: [Mandatory] To be called one time for initialising the library.
2) ConfigureThreshold: To be called for setting up any alerts
3) InitExpense: To be called along with any tags, before any metric value needs to be increased. 
        Based on the response from the library, client can choose to proceed or stop the execution.


Cache Refresh Strategies: The SDK maintains information about the alerts that are firing currently.
Strategy can be chosen while initialising the library.

We will be implementing multiple cache refresh strategies, the client would be able to choose the strategy based on use case.

Pull based strategy would be default. We may introduce more strategies in the future, keeping everything backward compatible.

Pull Based: SDK will pull the alert status and cache it.
Risk: Expense would not be stopped immediately, client will get to know about limit breach only after cache refresh.

Push Based[NOT YET IMPLEMENTED]: Client needs to implement a hook, a call will be made to this hook in case of any alert is firing or alert is recovered.
Challenge: Cascading the info to all the pods.


Testing:
Use main.go for any testing.