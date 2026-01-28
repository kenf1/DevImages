FROM rocker/rstudio:4.5.2

USER root
RUN apt-get update && apt-get install -y --no-install-recommends \
    libcurl4-openssl-dev \
    libssl-dev \
    libxml2-dev \
    libgit2-dev \
    && rm -rf /var/lib/apt/lists/*

USER rstudio
COPY --chown=rstudio:rstudio ./Dockerfiles/rc_ide.r /home/rstudio/rc_ide.r
RUN Rscript /home/rstudio/rc_ide.r

USER root
EXPOSE 8787
USER rstudio
