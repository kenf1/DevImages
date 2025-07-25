#base + tidyverse
FROM rocker/tidyverse:4.5.1

#rmarkdown specific
RUN apt-get update && apt-get install -y --no-install-recommends \
    wget graphviz lmodern perl && \ 
    /rocker_scripts/install_pandoc.sh && \
    install2.r rmarkdown

#task specific
RUN apt-get update && apt-get install -y \
    libcurl4-openssl-dev libssl-dev libxml2-dev \
    libharfbuzz-dev libfribidi-dev \
    libfontconfig1-dev libfreetype6-dev

COPY ./Dockerfiles/rc_rmd.R ./rc_rmd.R
RUN Rscript -e "source('./rc_rmd.R')"

ENV RSTUDIO_PANDOC=/usr/lib/rstudio/bin/pandoc