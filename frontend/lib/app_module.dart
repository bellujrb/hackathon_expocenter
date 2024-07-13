import 'package:flutter_modular/flutter_modular.dart';
import 'package:frontend/module/auth/presentation/screens/welcome_screen.dart';
import 'package:frontend/module/home/presentation/screens/home_screen.dart';
import 'package:frontend/module/home/presentation/screens/marketplace_screen.dart';
import 'package:frontend/module/home/presentation/screens/settings.dart';

class AppModule extends Module {

  @override
  List<Bind> get binds => [

  ];

  @override
  List<ChildRoute> get routes => [
    ChildRoute('/', child: (context, args) => const WelcomeScreen()),
    ChildRoute('/home', child: (context, args) => const HomeScreen()),
    ChildRoute('/exchange', child: (context, args) => const MarketPlaceScreen()),
    ChildRoute('/settings', child: (context, args) => const SettingsScreen())
  ];
}