import 'package:flutter/material.dart';
import 'package:flutter_modular/flutter_modular.dart';
import 'package:frontend/core/extensions/build_context_utils.dart';
import 'package:frontend/module/home/presentation/widgets/settings_menu_item.dart';
import 'package:google_fonts/google_fonts.dart';

class SettingsScreen extends StatefulWidget {
  const SettingsScreen({super.key});

  @override
  State<SettingsScreen> createState() => _SettingsScreenState();
}

class _SettingsScreenState extends State<SettingsScreen> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: null,
        toolbarHeight: 40,
        title: Padding(
          padding: const EdgeInsets.all(8.0),
          child: InkWell(
            onTap: () => Modular.to.navigate("profile"),
            child: SizedBox(
              height: 24,
              width: 24,
              child: Image.asset('icon_arrowleft.png'),
            ),
          ),
        ),
      ),
      backgroundColor: null,
      body: SizedBox(
        height: context.mediaHeight * 1.0,
        width: context.mediaWidth * 1.0,
        child: Column(
          children: [
            Text(
              "Settings",
              style: GoogleFonts.poppins(),
            ),
            const SizedBox(
              height: 10,
            ),
            SettingsMenuItem(
              title: "Notification",
              onTap: () {},
              img: "notifications-icon.png",
            ),
            SettingsMenuItem(
              title: "Theme",
              onTap: () {},
              img: "theme-icon.png",
            ),
            SettingsMenuItem(
              title: "Language",
              onTap: () {},
              img: "language-icon.png",
            ),
            SettingsMenuItem(
              title: "Rate our app",
              onTap: () {},
              img: "rate-icon.png",
            ),
          ],
        ),
      ),
    );
  }
}
