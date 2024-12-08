import 'dart:convert';
import 'dart:math' as math;

import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:intl/intl.dart';
import 'package:timeago/timeago.dart' as timeago;
import 'lat_lng.dart';
import 'place.dart';
import 'uploaded_file.dart';
import '/backend/backend.dart';
import 'package:cloud_firestore/cloud_firestore.dart';

LatLng generateCoordinate(int coordType) {
  // generate coordinates of astana, kabanbay batyr 53 if coordtype=1, else astana, turkestan 32
  if (coordType == 1) {
    return LatLng(51.128207, 71.430411);
  } else {
    return LatLng(51.180100, 71.445977);
  }
}
