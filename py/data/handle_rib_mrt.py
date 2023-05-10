import mrtparse
import yaml
# from .handlers import *

with open("py/ribs.yaml","w") as f:
    pass
for entry in mrtparse.Reader("test/testmrt/20230506.1551.table.dump"):
    pass
    with open("py/ribs.yaml","a") as f:
        yaml.dump([entry.data],f)