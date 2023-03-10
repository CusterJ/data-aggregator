package file_reader

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"reader/modules/domain"
)

type Entry struct {
	Error   error
	FooData domain.FooData
}

type Stream struct {
	stream chan Entry
}

func NewJsonStream() Stream {
	return Stream{
		stream: make(chan Entry), // try to remove "stream:"
	}
}

func (s Stream) Watch() <-chan Entry {
	return s.stream
}

func (s Stream) Start(path string) {
	defer close(s.stream)

	file, err := os.Open(path)
	if err != nil {
		s.stream <- Entry{
			Error: fmt.Errorf("open file: %w", err),
		}
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	if _, err := decoder.Token(); err != nil {
		s.stream <- Entry{
			Error: fmt.Errorf("decode opening delimiter: %w", err),
		}
		return
	}

	i := 0
	for decoder.More() {
		var fooData domain.FooData
		if err := decoder.Decode(&fooData); err != nil {
			s.stream <- Entry{
				Error: fmt.Errorf("decode line  %d: %w", i, err),
			}
			return
		}
		s.stream <- Entry{FooData: fooData}
		i++
	}

	// log.Println("fooData readed: ", i)

	if _, err := decoder.Token(); err != nil {
		s.stream <- Entry{
			Error: fmt.Errorf("decode closing delimiter: %w", err),
		}
		return
	}
}

func ReadFileWithStream(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	fmt.Println("func readFile with token: ", filename)
	d := json.NewDecoder(file)

	if _, err := d.Token(); err != nil {
		log.Println("decode Token error: ", err)
		return err
	}

	for d.More() {
		// s, _ := d.Token()
		var fd domain.FooData
		if err := d.Decode(&fd); err != nil {
			log.Fatal(err)
			return err
		}
		fmt.Printf("read %q\n", fd)
	}

	return nil
}

func ReadFullFile(filename string) ([]domain.FooData, error) {
	var data []domain.FooData

	dataString, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(dataString, &data)
	if err != nil {
		return nil, err
	}

	// log.Println("Len of file dataset is: ", len(data))
	return data, nil
}

func ReadFileWithBufTocken(filename string) error {
	var data domain.Dataset

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()
	fmt.Println("func readFile: ", filename)

	scanner := bufio.NewScanner(file)
	dataString := scanner.Text()
	fmt.Println("DataString is: ", string(dataString))

	json.Unmarshal([]byte(dataString), &data)

	fmt.Println("Len of dataset is: ", len(data.Dataset))
	fmt.Println("Data dataset is: ", data)

	return nil
}

func ReadFileWithBuf(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	fmt.Println("func readFile: ", filename)

	reader := bufio.NewReader(file)

	buf := make([]byte, 16)

	i := 1
	for {
		n, err := reader.Read(buf)

		if err != nil {

			if err != io.EOF {

				log.Fatal(err)
			}

			break
		}

		fmt.Print(string(buf[0:n]))
		i++
	}

	var data domain.Dataset
	fmt.Println(data)
	return nil
}

func ReadFileWithTokens(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	fmt.Println("func readFile with token: ", filename)
	d := json.NewDecoder(file)

	if _, err := d.Token(); err != nil {
		return err
	}

	for d.More() {
		// s, _ := d.Token()
		var fd domain.FooData
		if err := d.Decode(&fd); err != nil {
			log.Fatal(err)
			return err
		}
		fmt.Printf("read %q\n", fd)
	}

	return nil
}

func ReadFileByLines(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	fmt.Println("func readFileWithTokens with token: ", filename)
	d := json.NewDecoder(file)

	for {
		var fd domain.FooData
		if err := d.Decode(&fd); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s: %s\n", fd.ID, fd.Data, fd.Signal)
	}
	return nil
}
