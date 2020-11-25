package main

import (
	"testing"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	k8s "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"
	k8stest "k8s.io/client-go/testing"
)

const (
	succeed = "\u2713"
	failed  = "\u2717"
)

func Test_get_pod_using_fake_client(t *testing.T) {
	client := fake.NewSimpleClientset(&v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "fake",
			Namespace: "fake",
		},
	})

	t.Log("Fetch the pod by pod name using the client-go API ")
	{
		pod, err := client.CoreV1().Pods("fake").Get("fake", metav1.GetOptions{})
		t.Log("\tTest 0: checking the error code response")
		{
			if err != nil {
				t.Fatalf("\t%s\t get pod returned an unexpected error : %v", failed, err)
			}
			t.Logf("\t%s\tclient go has return no error.", succeed)
		}

		t.Log("\tTest 1: verifying the retrived pod from client-go get pod API")
		{
			if pod.Name != "fake" {
				t.Fatalf("\t%s\t got pod %s, expected %s", failed, pod.Name, "fake")
			}
			t.Logf("\t%s\tclient go has returned the expected pod", succeed)
		}
	}
}

func Test_watch_pod_using_fake_client(t *testing.T) {
	pods := []*v1.Pod{
		&v1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Labels: map[string]string{
					"app": "fake",
				},
				Name:      "fake",
				Namespace: "fake",
			},
			Status: v1.PodStatus{
				Phase: v1.PodPending,
			},
		},
		&v1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Labels: map[string]string{
					"app": "fake",
				},
				Name:      "fake",
				Namespace: "fake",
			},
			Status: v1.PodStatus{
				Phase: v1.PodUnknown,
			},
		},
		&v1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Labels: map[string]string{
					"app": "fake",
				},
				Name:      "fake",
				Namespace: "fake",
			},
			Status: v1.PodStatus{
				Phase: v1.PodRunning,
			},
		},
	}
	clients := simulatePodUpdates(pods)

	t.Log("Watch pod updates by pod name using the client-go API ")
	{
		watcher, err := clients.CoreV1().Pods("fake").Watch(metav1.ListOptions{
			LabelSelector: "app=fake",
		})

		t.Log("\tTest 0: checking the error code response")
		{
			if err != nil {
				t.Fatalf("\t%s\t watch pod returned an unexpected error : %v", failed, err)
			}
			t.Logf("\t%s\tclient go has return no error.", succeed)
		}

		t.Log("\tTest 1: checking watch event updates")
		{
			ch := watcher.ResultChan()
			for event := range ch {
				pod, ok := event.Object.(*v1.Pod)

				if !ok {
					t.Fatalf("\t%s\t watch event returned an unxpected error : %v", failed, ok)
				}
				t.Logf("\t%s\tgot a pod update event", succeed)

				if pod.Status.Phase == "" {
					t.Fatalf("\t%s\t expecting pod phase but its nil", failed)
				}
				t.Logf("\t%s\tgot a pod phase values %v", succeed, pod.Status.Phase)
			}
		}
	}
}

func simulatePodUpdates(pods []*v1.Pod) k8s.Interface {
	clients := fake.NewSimpleClientset()
	watcher := watch.NewFake()

	clients.PrependWatchReactor("pods", k8stest.DefaultWatchReactor(watcher, nil))

	go func() {
		defer watcher.Stop()

		for i, _ := range pods {
			time.Sleep(300 * time.Millisecond)
			watcher.Add(pods[i])
		}
	}()

	return clients
}
