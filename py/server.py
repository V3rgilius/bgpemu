from bokeh.layouts import column
from bokeh.models import ColumnDataSource, Slider
from bokeh.plotting import figure
from bokeh.sampledata.sea_surface_temperature import sea_surface_temperature
from bokeh.server.server import Server
from bokeh.plotting import figure
from bokeh.models import Select, Tabs
from bokeh.layouts import column
from bokeh.server.server import Server
from bgpdata.bokeh_docs import *
from bgpdata.handle_updates import *

def bkapp(doc):
    df = sea_surface_temperature.copy()
    source = ColumnDataSource(data=df)

    plot = figure(x_axis_type='datetime', y_range=(0, 25), y_axis_label='Temperature (Celsius)',
                  title="Sea Surface Temperature at 43.18, -70.43")
    plot.line('time', 'temperature', source=source)

    def callback(attr, old, new):
        if new == 0:
            data = df
        else:
            data = df.rolling(f"{new}D").mean()
        source.data = ColumnDataSource.from_df(data)

    slider = Slider(start=0, end=30, value=0, step=1, title="Smoothing by N Days")
    slider.on_change('value', callback)

    doc.add_root(column(slider, plot))

    # doc.theme = Theme(filename="theme.yaml")



server = Server({
    '/': get_topo_doc,
    '/spread': get_spread_doc,
    '/paths': get_path_change_doc,
    "/conv": get_convergence_doc,
    '/updates':get_updates_doc
    })
server.start()

if __name__ == '__main__':
    print('Opening Bokeh application on http://localhost:5006/')
    # server.io_loop.add_callback(server.show, "/")
    server.io_loop.start()

# paths,detail = get_path_change(analyzer.links,analyzer.elems,"6939","104.244.42.0/24")
# pass

# times = get_routing_convergence_times(analyzer.elems,3)

# cnts = get_packets_count(analyzer.elems,1,"136168")
# for cnt in cnts:
#     print(cnt)