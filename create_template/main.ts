//prompt user for input
const lang: string | null = prompt("Enter language name:");

if (lang == null) {
  console.error("Please enter a value");
} else {
  //file paths
  const dockerfile_path: string = `./Dockerfiles/${lang}Dev`;
  const workflow_path: string = `./.github/workflows/Build${lang}Dev.yml`;

  //workflow yaml contents
  const lang_lowercase: string = lang.toLowerCase();
  const workflow_template: string = `name: Build ${lang}Dev

on:
  workflow_dispatch:

jobs:
  build-and-push:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v4

      - name: Setup QEMU
        uses: docker/setup-qemu-action@v3

      - name: Setup Docker Buildx
        uses: docker/setup-buildx-action@v3

      - name: Login to GHCR
        uses: docker/login-action@v3
        with:
          registry: ghcr.io
          username: \$\{{ github.actor \}}
          password: \$\{{ secrets.GITHUB_TOKEN \}}

      - name: Build and push Docker image
        uses: docker/build-push-action@v6
        with:
          context: .
          file: ./Dockerfiles/${lang}Dev
          push: true
          platforms: linux/amd64
          tags: ghcr.io/kenf1/${lang_lowercase}dev:latest`;

  //create files
  try {
    await Deno.writeTextFile(workflow_path, workflow_template);
    await Deno.writeTextFile(dockerfile_path, "");

    console.log(`Files: ${workflow_path} & ${dockerfile_path} created.`);
  } catch (error) {
    console.error(`Error writing to file: ${error}`);
  }
}
