import 'package:flutter/material.dart';
import 'package:frontend/core/extensions/build_context_utils.dart';

class TransactionsWidget extends StatelessWidget {
  final String img;
  final String title;
  final String price;
  final String priceConvert;
  final String tokenConvert;
  final String date;
  const TransactionsWidget({
    super.key,
    required this.img,
    required this.title,
    required this.price,
    required this.priceConvert,
    required this.tokenConvert,
    required this.date,
  });

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      height: 72,
      width: context.mediaWidth * 1.0,
      child: Row(
        children: [
          Container(
            alignment: Alignment.center,
            height: 48,
            width: 48,
            decoration: BoxDecoration(
              borderRadius: BorderRadius.circular(100),
            ),
            child: SizedBox(
              height: 24,
              width: 24,
              child: Image.asset("assets/$img"),
            ),
          ),
          const SizedBox(
            width: 20,
          ),
          Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            mainAxisAlignment: MainAxisAlignment.spaceEvenly,
            children: [
              Text(
                title,
              ),
              Padding(
                padding: const EdgeInsets.only(right: 5),
                child: Row(
                  children: [
                    Text(
                      price,
                      style: const TextStyle(
                          color: Colors.black, fontSize: 16),
                    ),
                    const SizedBox(
                      width: 10,
                    ),
                    Row(
                      children: [
                        const Text(
                          "\$",
                          style:
                              TextStyle(color: Color(0xFFF8FAFC), fontSize: 16),
                        ),
                        Text(
                         "CASHBACK",
                          style: const TextStyle(
                              color: Color(0xFFF8FAFC), fontSize: 16),
                        ),
                      ],
                    ),
                    const SizedBox(
                      width: 10,
                    ),
                    SizedBox(
                      height: 16,
                      width: 16,
                      child: Image.asset("assets/$img"),
                    ),
                    const SizedBox(
                      width: 10,
                    ),
                    Row(
                      children: [
                        const Text(
                          "\$",
                          style:
                              TextStyle(color: Color(0xFFF8FAFC), fontSize: 16),
                        ),
                        Text(
                          tokenConvert,
                          style: const TextStyle(
                              color: Colors.red, fontSize: 16),
                        ),
                      ],
                    ),
                  ],
                ),
              )
            ],
          ),
          const SizedBox(
            width: 10,
          ),
          Text(
            date,
          )
        ],
      ),
    );
  }
}
