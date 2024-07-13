import 'package:flutter/material.dart';
import 'package:flutter_modular/flutter_modular.dart';
import 'package:frontend/core/styles/colors.dart';

AppBar buildAppBar(BuildContext context) {
  return AppBar(
    toolbarHeight: 80,
    title: InkWell(
      onTap: () => Modular.to.navigate("profile"),
      child: SizedBox(
        height: 40,
        width: 40,
        child: Image.asset("assets/profile.png"),
      ),
    ),
    actions: [
      Padding(
        padding: const EdgeInsets.symmetric(
          vertical: 20,
        ),
        child: Container(
          height: 40,
          width: 132,
          decoration: BoxDecoration(
            color: AppColors.primary,
            borderRadius: BorderRadius.circular(20),
          ),
          alignment: Alignment.center,
          child: const Text(
            "Convide e Ganhe",
            style: TextStyle(color: Colors.white),
          ),
        ),
      ),
      const SizedBox(
        width: 10,
      ),
      const Icon(
        Icons.notifications,
        color: Colors.black,
      ),
      const SizedBox(
        width: 10,
      ),
      const Icon(
        Icons.logout,
        color: Colors.black,
      ),
      const SizedBox(
        width: 10,
      ),
    ],
  );
}
