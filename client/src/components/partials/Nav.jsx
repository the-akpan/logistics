import { Link } from "react-router-dom";

import Container from "../layouts/Container";

function Nav() {
  return (
    <nav className=" text-white bg-primary py-5 drop-shadow-md">
      <Container>
        <div className="flex justify-between items-center">
          <Link to="/">
            <h5 className="text-2xl font-bold">B<span className="text-orange-500">T</span>L<span className="">ogistics</span></h5>
          </Link>
        </div>
      </Container>
    </nav>
  );
}

export default Nav;
