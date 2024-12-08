// Automatic FlutterFlow imports
import '/backend/backend.dart';
import '/flutter_flow/flutter_flow_theme.dart';
import '/flutter_flow/flutter_flow_util.dart';
import 'index.dart'; // Imports other custom actions
import '/flutter_flow/custom_functions.dart'; // Imports custom functions
import 'package:flutter/material.dart';
// Begin custom action code
// DO NOT REMOVE OR MODIFY THE CODE ABOVE!

import 'package:map_launcher/map_launcher.dart';

Future mapLauncher(LatLng coords) async {
  final availableMaps = await MapLauncher.installedMaps;
  print(
      availableMaps); // [AvailableMap { mapName: Google Maps, mapType: google }, ...]

  await availableMaps.first.showMarker(
    coords: Coords(coords.latitude, coords.longitude),
    title: "end location",
  );
  // Add your function code here!
}
