include .env
export


front-run:
	cd ./frontend && npm run dev

pull-ui:
	cd ./frontend && npm run add-component $(FRONT_SHAD_UI_COMP)