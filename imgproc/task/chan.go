package task

import (
	"fmt"
	"path"
	"path/filepath"

	"training.go/imgproc/filter"
)

type ChanTask struct {
	dirCtx
	Filter   filter.Filter
	PoolSize int
}

func NewChanTask(srcDir string, dstDir string, filter filter.Filter, poolSize int) Tasker {

	return &ChanTask{
		Filter: filter,
		dirCtx: dirCtx{
			SrcDir: srcDir,
			DstDir: dstDir,
			Files:  BuildFileList(srcDir),
		},
	}
}

type jobReq struct {
	src string
	dst string
}

func (c *ChanTask) Process() error {
	size := len(c.Files)
	jobs := make(chan jobReq, size)
	results := make(chan string, size)

	// init workers
	for w := 1; w <= c.PoolSize; w++ {
		go worker(w, c, jobs, results)
	}

	// Start jobs
	for _, f := range c.Files {
		filename := filepath.Base(f)
		dst := path.Join(c.DstDir, filename)
		jobs <- jobReq{
			src: f,
			dst: dst,
		}
	}

	close(jobs)

	for range c.Files {
		fmt.Println(<-results)
	}

	return nil
}

func worker(id int, chanTask *ChanTask, jobs <-chan jobReq, results chan<- string) {
	for j := range jobs {
		fmt.Printf("Worker %d, started job %v\n", id, j)
		chanTask.Filter.Process(j.src, j.dst)
		results <- j.dst
	}
}
