
from flask import render_template, jsonify
from app import app

@app.route('/')
def index():
    return render_template('index.html')


@app.route('/about')
def about():
    return render_template('about.html')

n = {"1":(8, 60), "2":(10,65)}
nodes = []
for k, x in n.items():
    nodes.append("""{
        "type": "Feature",
        "properties": {},
        "geometry": {
            "type": "Point",
            "coordinates": [
            """+ str(x[0]) + """,
            """+ str(x[1]) + """
            ]
        }
        }""")


point = '''{
  "type": "FeatureCollection",
  "features": [
    ''' + ",".join(nodes) + '''
  ]
}'''

@app.route('/map')
def map():
    return render_template('map.html', routeString=point)

