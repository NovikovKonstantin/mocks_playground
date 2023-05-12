PLAYGROUND_DIR=playground
TEMPLATE_DIR=template
PROJECT_NAME=$(name)
PROJECT_DIR=$(PLAYGROUND_DIR)/$(PROJECT_NAME)

# Create a new mockable service in the playground.
create:
ifeq ($(PROJECT_NAME),)
	@echo Name parameter is empty.
	@echo Provide a name like this: 'make create name=name-of-the-project'.
	@exit 1
endif
	@echo Trying to create '$(PROJECT_NAME)' project...

	@if [ -d $(PROJECT_DIR) ]; then \
		echo Project $(PROJECT_NAME) is already exists; \
		echo Check $(PROJECT_DIR) path; \
		exit 1; \
	fi

	@cp -R $(TEMPLATE_DIR) $(PROJECT_DIR)

	@( cd $(PROJECT_DIR) && go mod init $(PROJECT_NAME) && go mod tidy )

	@echo Done
