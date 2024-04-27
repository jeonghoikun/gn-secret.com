.PHONY:
	tailwindc

tailwindc:
	npx tailwindcss -i ./static/css/tailwindcss.css -o ./static/css/main.css --watch
