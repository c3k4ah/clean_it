package contents

const DtoModelTemplate = `
class {{.Model}} extends {{.Entity}} {
  // Define your model fields here

  {{.Model}}({
    // Initialize your fields here
  });

  factory {{.Model}}.fromJson(Map<String, dynamic> json) {
    return {{.Model}}(
      // Parse your JSON fields here
    );
  }

  Map<String, dynamic> toJson() {
    return {
      // Convert your fields to JSON here
    };
  }
}
`

const DtoEntityTemplate = `
import 'package:equatable/equatable.dart';

class {{.Entity}} extends Equatable {
  // Define your entity fields here

  {{.Entity}}({
    // Initialize your fields here
  });

  @override
  List<Object> get props => [
    // Return your fields here
  ];
}
`
