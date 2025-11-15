TARGET := trident-agent

.PHONY: all $(TARGET) clean

all: $(TARGET)

$(TARGET):
	@echo "\033[1;36mBuilding $@...\033[1;37m"
	go build -o ./build/$@ ./cmd/$@

clean:
	rm -rf build/
