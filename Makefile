BRANCH="develop"

help:
	echo 'make init: init remote urls'
	echo 'make pull: pull all project'
	echo 'make push: push all project'

init:
	git remote add core git@github.com:ochanoco/ochano.co-core.git
	git remote add auth git@github.com:ochanoco/ochano.co-auth.git
	git remote add cloud git@github.com:ochanoco/ochano.co-cloud.git
	git remote add tee git@github.com:ochanoco/ochano.co-tee.git

pull:
	git subtree pull --prefix=core core ${BRANCH}
	git subtree pull --prefix=auth auth ${BRANCH}
	git subtree pull --prefix=cloud cloud ${BRANCH}
	git subtree pull --prefix=tee tee ${BRANCH}


push:
	git subtree push --prefix=core core ${BRANCH}
	git subtree push --prefix=auth auth ${BRANCH}
	git subtree push --prefix=cloud cloud ${BRANCH}
	git subtree push --prefix=tee tee ${BRANCH}
