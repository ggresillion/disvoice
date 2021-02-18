#pragma once

#include <stdexcept>
#include <iostream>

using namespace std;

class Error : std::exception
{
public:
	Error(const std::string &message) : message(message)
	{
		cout << message << endl;
	}

	const std::string &getMessage() const
	{
		return message;
	}

	std::string message;
};
