publishpolicy:
	pip3 install -r policies/policypack/requirements.txt
	cd policies/policypack && pulumi policy publish testmarrod
	pulumi policy enable testmarrod/aws-python 0.0.1
