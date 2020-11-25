
# Use the fake watch API and simulate the watch behaviour sequence


Sometimes we encounter the case where we need to simulate the watch events in 
order to test code that uses client-go `watch` API.

In this example, we will see how to mock `watch` API events sequence as part of the unit testing.

Before that, let's see how to use the `fake` package to `client-go` APIs.

In this case, we will use the `pod()` API.

## Test get pod by name

Following snippet initialise the fake client by feeding a  `pod` resource.
```
import(
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

client := fake.NewSimpleClientset(&v1.Pod{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "fake",
		Namespace: "fake",
	},
})
```

Test it by running `make test-get-pod`

```
➜  mock-watch-events git:(mock-watch) ✗ make test-get-pod
go test -run Test_get_pod_using_fake_client -v
=== RUN   Test_get_pod_using_fake_client
    fake_client_test.go:28: Fetch the pod by pod name using the client-go API 
    fake_client_test.go:31: 	Test 0: checking the error code response
    fake_client_test.go:36: 	✓	client go has return no error.
    fake_client_test.go:39: 	Test 1: verifying the retrived pod from client-go get pod API
    fake_client_test.go:44: 	✓	client go has returned the expected pod
--- PASS: Test_get_pod_using_fake_client (0.00s)
PASS
ok  	github.com/hrishin/k8s-client-go-examples/examples/mock-watch-events	0.571s
```

# Test pod watch events mocking

In this senario, we will simulate pod life cycle events i.e. `pod.status.phase` -> {PodPending,  PodUnknown, PodRunning}.
Usually, we encounter such code to wait for the pod to become up & running.

To feed such events `client-go` provide `testing` package. Following example snippet the way to feed the mock events for the `watch` API.

```
clients := fake.NewSimpleClientset()
watcher := watch.NewFake()
clients.PrependWatchReactor("pods", k8stest.DefaultWatchReactor(watcher, nil))

go func() {
	defer watcher.Stop()

	for i, _ := range pods {
		time.Sleep(300 * time.Millisecond)
		watcher.Add(&v1.Pod{
		   ..... // your pod definitions
		})
	}
}()
```

Important to note here that in `clients.PrependWatchReactor("pods", k8stest.DefaultWatchReactor(watcher, nil))`  method `pods` is the plural resource name. Giving the wrong resource name would fail mocking watch events. One of the way to get the resource name is using 
```
kubectl api-resources | grep -h pod

NAME  SHORTNAMES APIGROUP NAMESPACED KIND
pod  po  				  true 		 Pod
```

Test the example by running `make test-watch-pod`

```
➜  mock-watch-events git:(mock-watch) ✗ make test-watch-pod
go test -run Test_watch_pod_using_fake_client -v
=== RUN   Test_watch_pod_using_fake_client
    fake_client_test.go:90: Watch pod updates by pod name using the client-go API 
    fake_client_test.go:96: 	Test 0: checking the error code response
    fake_client_test.go:101: 	✓	client go has return no error.
    fake_client_test.go:104: 	Test 1: checking watch event updates
    fake_client_test.go:113: 	✓	got a pod update event
    fake_client_test.go:118: 	✓	got a pod phase values Pending
    fake_client_test.go:113: 	✓	got a pod update event
    fake_client_test.go:118: 	✓	got a pod phase values Unknown
    fake_client_test.go:113: 	✓	got a pod update event
    fake_client_test.go:118: 	✓	got a pod phase values Running
--- PASS: Test_watch_pod_using_fake_client (0.90s)
PASS
ok  	github.com/hrishin/k8s-client-go-examples/examples/mock-watch-events	4.468s
```

I hope this post will be useful. Would like to hear your reviews, feedback or your experience. 
Happy programming with Kubernetes!