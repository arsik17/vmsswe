// Automatic FlutterFlow imports
import '/backend/backend.dart';
import '/flutter_flow/flutter_flow_theme.dart';
import '/flutter_flow/flutter_flow_util.dart';
import 'index.dart'; // Imports other custom widgets
import '/custom_code/actions/index.dart'; // Imports custom actions
import '/flutter_flow/custom_functions.dart'; // Imports custom functions
import 'package:flutter/material.dart';
// Begin custom widget code
// DO NOT REMOVE OR MODIFY THE CODE ABOVE!

import 'dart:math';
import 'package:google_maps_widget/google_maps_widget.dart' as GM;

class GMap extends StatefulWidget {
  const GMap({
    Key? key,
    this.width,
    this.height,
  }) : super(key: key);

  final double? width;
  final double? height;

  @override
  _GMapState createState() => _GMapState();
}

class _GMapState extends State<GMap> {
  final mapsWidgetController = GlobalKey<GM.GoogleMapsWidgetState>();

  @override
  Widget build(BuildContext context) {
    return GM.GoogleMapsWidget(
      apiKey: "AokcpkogkokKDCIOJojpojC9032-_KCMDi39",
      key: mapsWidgetController,
      sourceLatLng: GM.LatLng(51.098779, 71.421663),
      destinationLatLng: GM.LatLng(51.088956, 71.397090),
      routeWidth: 3,
      destinationMarkerIconInfo: GM.MarkerIconInfo(
        assetPath: "assets/images/end-marker-icon.png",
        // assetMarkerSize: Size.square(70),
      ),
      driverMarkerIconInfo: GM.MarkerIconInfo(
        infoWindowTitle: "Alex",
        assetPath: "assets/images/driver-marker-icon.png",
        assetMarkerSize: Size.square(70),
      ),
    );
  }
}
