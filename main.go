package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Create a span for a web request at the /posts URL.
	timenow := time.Now().Unix()
	// tracer.StartSpanFromContext(ctx, "xxx")
	span2, ctx2 := tracer.StartSpanFromContext(r.Context(), fmt.Sprintf("web.request %v", timenow))
	log.Printf("my log messagezzz  %v", span2)

	spandID := span2.Context().SpanID()
	fmt.Printf("spanid: %v\n", spandID)
	//fmt.Println("ctx2: ", ctx2)
	defer span2.Finish()
	//ctx4 := context.WithValue(ctx2, "abc", "def")

	// Set tag
	span2.SetTag("http.url222", r.URL.Path)
	span2.SetTag("test_tag222", "tag_value")
	// ctx3 := tracer.ContextWithSpan(ctx2, span2)

	TestTrace(ctx2, timenow)
	// TestTrace2(ctx4, timenow)

}

func TestTrace(ctx context.Context, timenow int64) {
	//span, ctx2 := tracer.StartSpanFromContext(ctx, " child 1 TestTrace1")
	span := tracer.StartSpan("child 1 TestTrace", tracer.ServiceName("hbc_test_service_name"))
	// if ok {
	spandID := span.Context().SpanID()
	fmt.Printf("spanid: %v\n", spandID)
	log.Printf("my TestTrace messagezzz  %v", span)

	fmt.Printf("\n\n\n======hello world --- %v ---- =======\n\n", timenow)
	defer span.Finish()
	time.Sleep(time.Second)

	// Set tag
	span.SetTag("test_tag", "tag_value")
	// }
	TestTrace2(ctx, timenow)
}

func TestTrace2(ctx context.Context, timenow int64) {
	span, _ := tracer.StartSpanFromContext(ctx, " child 2 TestTrace2")
	//span, _ := tracer.StartSpanFromContext(ctx, " 22 child span TestTrace 3333")
	//span := tracer.StartSpan("hello 22222 ")
	spandID := span.Context().SpanID()
	fmt.Printf("spanid: %v\n", spandID)
	time.Sleep(time.Second)
	data := ctx.Value("abc")
	fmt.Printf("data: %v \n\n", data)
	ctx5 := context.WithValue(ctx, "abc1", "xxxx")
	data1 := ctx5.Value("abc1")
	fmt.Printf("data: %v \n\n", data1)

	data2 := ctx.Value("abc1")
	fmt.Printf("data: %v \n\n", data2)

	// if ok {
	fmt.Printf("\n\n\n======hello world --- %v ---- =======\n\n", timenow)
	defer span.Finish()
	// Set tag
	span.SetTag("test_tag", "tag_value")
	// }
}

func main() {
	tracer.Start(tracer.WithService("hbc_test_service_name"))
	defer tracer.Stop()
	http.HandleFunc("/posts", handler)
	log.Fatal(http.ListenAndServe(":4562", nil))
}

//DD_AGENT_MAJOR_VERSION=7 DD_API_KEY=9b5d46b4e41e2bc67360889133f75665 DD_SITE="us5.datadoghq.com" bash -c "$(curl -L https://s3.amazonaws.com/dd-agent/scripts/install_mac_os.sh)"
