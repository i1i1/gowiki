GOWIKI=gowiki
MD_PAGES=$(shell find md -name '*.md' -type f)
HTML_PAGES=$(subst md, html, $(MD_PAGES:.md=.html)) html/index.html

all: $(HTML_PAGES) httpd.conf .dcbuild

.dcbuild:
	touch .dcbuild
	docker build -t i1i1/wiki .

run: all
	docker run -d -v $(shell pwd)/html:/usr/local/apache2/htdocs/ -p 80:80 i1i1/wiki

html/%.html: md/%.md
	$(GOWIKI) <$< >$@

md/index.md: ./generindex $(filter-out md/index.md, $(MD_PAGES))
	./generindex $(MD_PAGES) >$@

clean:
	rm -f $(HTML_PAGES) md/index.md .dcbuild

.PHONY: all run clean
