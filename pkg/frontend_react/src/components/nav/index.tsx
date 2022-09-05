import * as React from 'react';
import { Nav, NavProps, PageSidebar } from '@patternfly/react-core';

import PerspectiveNav from './PerspectiveNav';
import NavHeader from './NavHeader';

type NavigationProps = {
  onNavSelect: NavProps['onSelect'];
  onPerspectiveSelected: () => void;
  isNavOpen: boolean;
};

export const Navigation: React.FC<NavigationProps> = React.memo(function Navigation({
  isNavOpen,
  onNavSelect,
  onPerspectiveSelected,
}) {
  return (
    <PageSidebar
      nav={
        <Nav aria-label={'Nav'} onSelect={onNavSelect} theme="dark">
          <NavHeader onPerspectiveSelected={onPerspectiveSelected} />
          <PerspectiveNav />
        </Nav>
      }
      isNavOpen={isNavOpen}
      theme="dark"
    />
  );
});
