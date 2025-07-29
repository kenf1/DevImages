import os


class Configs:
    cwd: str = os.getcwd()

    dev_req: str = "dev_requirements.txt"
    dev_full_path: str = f"{cwd}/setup/{dev_req}"

    prod_req: str = "prod_requirements.txt"
    prod_full_path: str = f"{cwd}/setup/{prod_req}"


def fix_requirements(to_rm: list):
    with open(Configs.dev_full_path, "r", encoding="utf-8") as file:
        lines = file.readlines()
        res = [x for x in lines if x not in to_rm]

    with open(Configs.prod_full_path, "w", encoding="utf-8") as f:
        f.write("".join(res))


if __name__ == "__main__":
    to_rm = ["ruff\n", "pytest\n", "pytest-asyncio\n", "\n"]
    fix_requirements(to_rm)
