FROM mcr.microsoft.com/dotnet/sdk:10.0

RUN apt-get update && apt-get install -y git make
ENV DOTNET_CLI_TELEMETRY_OPTOUT=1