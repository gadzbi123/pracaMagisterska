:
f=main
#minted=
minted=-shell-escape 
minteddir=_minted-$(f)

.PHONY : short
.PHONY : clean

setup:
	evince main.pdf &
	nvim .

xelatex : $(f).tex
	xelatex $(minted) -synctex=1 $(f)
	bibtex $(f)
	xelatex $(minted) -synctex=1 $(f)
	xelatex $(minted) -synctex=1 $(f)

pdflatex : $(f).tex
	pdflatex $(minted) -synctex=1 $(f)
	bibtex $(f)
	pdflatex $(minted) -synctex=1 $(f)
	pdflatex $(minted) -synctex=1 $(f)

$(f).pdf : $(f).tex
	pdflatex $(minted) -synctex=1 $(f)
	bibtex $(f)
	pdflatex $(minted) -synctex=1 $(f)
	pdflatex $(minted) -synctex=1 $(f)

build : $(f).tex
	pdflatex $(minted) -synctex=1 $(f)
	bibtex $(f)
	pdflatex $(minted) -synctex=1 $(f)
	pdflatex $(minted) -synctex=1 $(f)

short : $(f).tex
	pdflatex $(minted) -synctex=1 $(f)

clean : 
	for p in *.aux *.bbl *.bcf *.blg *.lof *.log *.lot *.out *.toc *.xml *.gz *-blx.bib main.pdf; \
	do \
		if [ -f $$p ] ; then rm $$p; fi; \
	done;  \
	if [ -d $(minteddir) ] ; then rm -r $(minteddir) ; fi; 
