include $(GOROOT)/src/Make.inc

EXE=wordfinder
OBJ=wordfinder.$(O) 
GOFMT=gofmt -spaces=true -tabindent=false -tabwidth=4
SRC=wordfinder.go

all: clean
	$(GC) -o ${OBJ} ${SRC}
	$(LD) -o ${EXE} ${OBJ}

run: all
	./${EXE}

clean:
	rm -f ${OBJ} ${EXE}

format:
	${GOFMT} -w ${SRC}

