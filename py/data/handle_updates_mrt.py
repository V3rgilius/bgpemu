import mrtparse
import yaml
# from .handlers import *


for entry in mrtparse.Reader("test/testmrt/20230506.1551.table.dump"):
    pass
    with open("py/updates.yaml","a") as f:
        yaml.dump([entry.data],f)