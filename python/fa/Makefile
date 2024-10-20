ROOT=python/fa/

.PHONY: devel
devel: reqs.txt 
	pip install reqs.txt
	#if ! poetry lock --check &> /dev/null; then \
	#	poetry lock; \
	#fi
	#poetry install


.PHONY: run
run:
	fastapi dev main.py


.PHONY: run_uv4
run_uv4:
	@uvicorn main:app --workers 4 --no-access-log --fd 0


.PHONY: run_uv8
run_uv8:
	uvicorn main:app --workers 8 --no-access-log


.PHONY: run_gc
run_gc:
	gunicorn main:app -w 4 -k uvicorn.workers.UvicornWorker


.PHONY: run_gc_h11
run_gc_h11:
	gunicorn main:app -w 4 -k uvicorn.workers.UvicornH11Worker


.PHONY: run_hc
run_hc:
	hypercorn main:app -w 4


.PHONY: run_gr
run_gr:
	granian --interface asgi main:app --workers 4


.PHONY: ab
ab:
	@ab -n 10000 -c 100 http://127.0.0.1:8000/hello
